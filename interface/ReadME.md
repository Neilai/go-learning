## Go的接口

- go的接口是鸭子类型，弱约束,兼具python和Java的优点
- 比较奇怪的一点是Go interface若是pointer调用，只能传指针，若是value 调用，传指针和值都可以
- interface可通过 x.type判断实际调用类型,见demo