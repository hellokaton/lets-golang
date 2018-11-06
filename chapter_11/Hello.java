public class Hello {

    private static void change(Integer num) {
        num = 2333;
        System.out.println(num);
    }

    public static void main(String[] args) {
        Integer num = 1;
        change(num);
        System.out.println(num);
    }

}