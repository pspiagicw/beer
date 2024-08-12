class Student {
    fn init(name string, surname string) {
        let this.name = name;
        let this.surname = surname;
    }
    fn hello() {
        print("Hello, " + this.name);
    }
    fn getName(name) string {
        return this.name + " " +this.surname;
    }
}

fn main() {
    v = Student("Chris", "Pratt");

    v.hello();
}
