import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        System.out.println("Введите число!");
        Scanner scanner = new Scanner(System.in);

        int num = scanner.nextInt();

        List<String> numbers = new ArrayList<>();

        for (int i = 0; i < num; i++) {
            if (getSimple(i)) {
                numbers.add(Integer.toString(i));
            }
        }

        File file = new File("Output.txt");
        try {
            file.createNewFile();
            FileOutputStream oFile = new FileOutputStream(file, true);
            oFile.write((numbers + "\n").getBytes());
        } catch (java.io.IOException e) {
            System.out.println(e);
        }
    }

    public static boolean getSimple(int i) {
        if (i <= 1) {
            return false;
        } else if (i <= 3) {
            return true;
        } else if (i%2 == 0 || i%3 == 0) {
            return false;
        }
        int n = 5;
        while (n*n <= i) {
            if (i%n == 0 || i%(n+2) == 0) {
                return false;
            }
            n = n + 6;
        }
        return true;
    }
}