package jssc;//import jssc.*;
import java.awt.*;
import java.awt.datatransfer.Clipboard;
import java.awt.datatransfer.StringSelection;
import java.awt.event.KeyEvent;
import java.io.IOException;
import java.util.ArrayList;

/**
 * Created by TrashPony on 15.11.16.
 */

public class weightDriver {
    private static SerialPort serialPort;
    private static ArrayList<Integer> dataBytes = new ArrayList<>();
    private static double result;
    private static double controlResult;
    private static String data = "";

    public static void main(String[] args) throws IOException {
        try {
            selectPort();
            while(true) {
                commands();
                if(data.isEmpty()) selectPort();
                data = "";
            }
        } catch (Exception e) {
            System.out.println(e);
        }
    }
    private static void selectPort() throws IOException {

       try {
           String[] portNames = SerialPortList.getPortNames();
            if(portNames.length == 0) {
            System.out.println("Ports not found!");
                Thread.sleep(3000);
                main(portNames);
            } else {
                for (int i = 0; i < portNames.length; i++) {
                    serialPort = new SerialPort(portNames[i]);
                    serialPort.openPort();
                    serialPort.setParams(SerialPort.BAUDRATE_4800,
                            SerialPort.DATABITS_8,
                            SerialPort.STOPBITS_1,
                            SerialPort.PARITY_EVEN);
                    serialPort.addEventListener(new PortReader(), SerialPort.MASK_RXCHAR);
                    commands();
                    Thread.sleep(100); // ожидаем ответа от весов
                    if (!(data.isEmpty())) {
                        System.out.println("Весы подключены к порту " + portNames[i]);
                        break;
                    }
                    if (data.isEmpty()) serialPort.closePort();
                    if (i == (portNames.length - 1)) {
                        Thread.sleep(3000);
                        System.out.println("Weight not found!");
                        i = -1;
                    }
                }
            }
        } catch (Exception e){
            System.out.println(e);
        }
    }

    private static void commands() throws IOException {
        try {

            // не/готовность 0/128 и дискретность 0х00-1г,0х01-0.1г,0х04-0.01кг,0.05-0.1кг
            serialPort.writeByte((byte) 0x48);
            Thread.sleep(50);
            //вес в виде 2х байтов n х n
            serialPort.writeByte((byte) 0x45);
            Thread.sleep(150);
            dataBytes.clear();

        } catch (Exception e) {
            selectPort();
        }
    }

    private static class PortReader implements SerialPortEventListener {
        private static String toClipBoard;
        public void serialEvent(SerialPortEvent event) {
            if (event.isRXCHAR() && event.getEventValue() > 0) {
                try {
                    data = serialPort.readHexString();
                    if(data.length() <= 2) {
                        int hexToBin = (Integer.parseInt(data, 16));
                        if (dataBytes.size() < 4) dataBytes.add(hexToBin);
                    }

                    if(dataBytes.size() == 4 && dataBytes.get(0) == 128 && (dataBytes.get(2) !=0 || dataBytes.get(3) !=0)){
                        if(dataBytes.get(1) == 4 && dataBytes.get(3) == 0) result = ((double)dataBytes.get(2)*0.01);
                        if(dataBytes.get(1) == 4 && dataBytes.get(3) != 0) result = ((double)((256*dataBytes.get(3))+dataBytes.get(2))*0.01);
                        if (!((controlResult -0.01) <= result && result <= (controlResult + 0.1))) {
                            toClipBoard = String.format("%.2f", result);
                            putToClipBoard(toClipBoard);
                            ctrlVEnter();
                            controlResult = result;
                        }
                    }

                    if(dataBytes.size() == 4 && dataBytes.get(0) == 0) {
                        controlResult = 0;
                    }
                } catch (Exception e) {


                }
            }
        }
    }

    // Копирование ответа весов в буффер обмена
    private static void putToClipBoard(String data) {
        Clipboard clbrd = Toolkit.getDefaultToolkit().getSystemClipboard();
        StringSelection strsel = new StringSelection(data);
        clbrd.setContents(strsel, null);
    }

    // Имитация нажатия клавиш CTRL+V+Enter
    private static void ctrlVEnter() {
        try {
            Robot robot = new Robot();
            robot.keyPress(17);     // нажимает ктрл
            robot.keyPress('V');    // нажимаем кнопу 'V' с клавиатуры
            robot.delay(20);
            robot.keyRelease('V');  // отжимаем кнопу 'V' с клавиатур
            robot.keyRelease(17);   // отжимаем ктрл
            robot.delay(20);
            robot.keyPress(KeyEvent.VK_ENTER);
            robot.delay(10);
            robot.keyRelease(KeyEvent.VK_ENTER);
            robot.delay(20);

        } catch (AWTException e) {
            System.err.println(e);
        }
    }
}
