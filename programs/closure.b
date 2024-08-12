fn countit() fn() int {
    let counter = 1;
    return fn() {
        let counter = counter + 1;
        return counter;
    };
}

fn main() {
    let c = countit();

    print(c());
    print(c());
    print(c());
    print(c());
}
