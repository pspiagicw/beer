void main() {
    def number int;
    let number = 1;

    if number == 1 {
        print("Always true");
    } else {
        print("Shouldn't execute");
    }

    if number == 2 {
            print("Number is 2");
    } else {
        print("Number is not 2");
    }


    for number = 1; number < 10; number++ {
        printf("Number is %s",number);
    }

    while number < 10 {
        print("Not infinity");
        let number = number + 1;
    }
}
