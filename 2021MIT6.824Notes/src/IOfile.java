import java.io.*;

public class IOfile {

    public static void main(String[] args) throws Exception {


        File file = new File("D:\\CS-openCourse\\2021MIT6.824Notes\\src\\test.txt");

        Thread thread1 = new Thread(()->{
            String context = null;
            BufferedReader bufferedReader = null;
            try {
                bufferedReader = new BufferedReader(new FileReader(file));
            } catch (FileNotFoundException e) {
                e.printStackTrace();
            }
            while (true){
                try {
                    if (!((context = bufferedReader.readLine()) != null)) break;
                } catch (IOException e) {
                    e.printStackTrace();
                }
                System.out.println(context);
            }
        });

        Thread thread2 = new Thread(()->{
            String context = "Test MIT6.824";
        });

        thread1.run();
    }
}
