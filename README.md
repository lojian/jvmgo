# jvmgo
A JVM written in GO programming language

This is part of my plan of learning Go programming language.

## build instructions
build is supported on Windows, Mac, and most popular Linux system.
Java 1.7 or later is required.

```bash
git clone https://github.com/lojian/jvmgo.git
cd jvmgo/jvmgo
./mvnw clean package
```
The generate binary is located at bin/javago
Note: first time execution will take some time to download GO SDK and Maven

## Java Testing Program
Following Java program could be found under javaclasses directory
```java
public class GaussTest {
    public static void main(String[] args) {
        int sum = 0;
        for (int i = 1; i <=100; i++) {
            sum += i;
        }
        System.out.println(sum);
    }
}
```

## VM execution demo
You need to setup JAVA_HOME
```
$ bin/javago -Xjre $JAVA_HOME/jre -cp ~/jvmgo/jvmgo/javaclasses/build/classes/java/main GaussTest
found method:<init>,  ()V
found method:main,  ([Ljava/lang/String;)V
pc:  0  inst: *constants.ICONST_0 &{{}}
pc:  1  inst: *stores.ISTORE_1 &{{}}
pc:  2  inst: *constants.ICONST_1 &{{}}
pc:  3  inst: *stores.ISTORE_2 &{{}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
pc: 10  inst: *loads.ILOAD_1 &{{}}
pc: 11  inst: *loads.ILOAD_2 &{{}}
pc: 12  inst: *math.IADD &{{}}
pc: 13  inst: *stores.ISTORE_1 &{{}}
pc: 14  inst: *math.IINC &{2 1}
pc: 17  inst: *control.GOTO &{{-13}}
pc:  4  inst: *loads.ILOAD_2 &{{}}
pc:  5  inst: *constants.BIPUSH &{100}
pc:  7  inst: *comparisons.IF_ICMPGT &{{13}}
LocalVars: [{0 <nil>} {5050 <nil>} {101 <nil>}]
OperandStack:&{0 [{101 <nil>} {100 <nil>}]}
```

you could see, the VM could calculate out 5050
