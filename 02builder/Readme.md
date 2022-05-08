## 建造者模式定义
- Separate the construction of a complex object from its representation so that the same construction process can create different representations.
- 将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

## 建造者模式的有点
- 封装性好
    1. 调用者 不需要关注 产品类的实现细节

## 建造者模式应用场景
- 初始化一个对象特别复杂，比如提供了多个构造方法，或者构造函数有很多参数，但创建对象时，并不想马上给成员属性赋值。
- 零件装配顺序不同，产生的对象 效果不同。