## Project 0: 2048

### project link

[Project 0: 2048 | CS 61B Spring 2021 (datastructur.es)](https://sp21.datastructur.es/materials/proj/proj0/proj0)

## project Introduction

熟悉Java以及 IntelliJ IDE 和 JUnit，project0的任务只在Model.java中，并且仅限于四个方法。

## The game

关于2048，你可能玩过，这是一款单人电脑游戏，由 Gabriele Cirulli 编写，基于 Veewo Studio 早期的游戏《1024》。

在这个项目中，你需要构造这个游戏的核心逻辑，也就是说，，我们已经将所有 GUI 代码、处理按键和大量其他脚手架放在一起了。你的工作将是做最重要和最有趣的部分。

具体来说，你将在 Model.java 文件中填写4个方法，这些方法管理用户按下某些键之后发生的事情。

这个游戏本身是十分简单的，在一个4 x 4的正方形网格上进行的，每个正方形可以是空的，也可以包含一个整数-大于或等于2的方块。在第一次移动之前，应用程序将一个包含2或4的瓦片添加到最初空白板上的随机方块中。2或4的选择是随机的，有75% 的几率选择2,25% 的几率选择4。

然后玩家通过他们的箭头键选择一个方向来倾斜棋盘: 北，南，东，西。所有的瓷砖都沿着这个方向滑动，直到在运动的方向上没有留下空隙(可能没有空隙可以开始)。一个瓷砖可能会与另一个瓷砖合并，从而赢得玩家点数。

## game rules

- 两个具有相同值的图像块合并成一个图像块，其中包含的数字是初始数字的两倍。
- 合并后的平铺图案在倾斜时不会再次合并。例如，如果我们有[ x，2,2,4] ，其中 x 代表一个空白区域，我们把这些图像移到左边，我们应该得到[4,4，x，x ] ，而不是[8，x，x，x，x ]。这是因为最左边的4已经是合并的一部分，所以不应该再次合并。
- 当运动方向上的三个相邻瓦片具有相同的数目时，运动方向上的两个前导瓦片合并，而拖曳瓦片不合并。例如，如果我们有[ x，2,2,2]并且向左移动图块，我们应该以[4,2，x，x ]而不是[2,4，x，x ]结束。
- 作为这些规则的必然结果，如果在运动方向上有四个相邻的瓦片具有相同的数目，它们就会形成两个合并的瓦片。例如，如果我们有[4,4,4,4,4] ，那么如果我们向左移动，我们最终得到[8,8，x，x ]。这是因为前两个瓷砖将根据规则3进行合并，然后后两个瓷砖将进行合并，但是根据规则2，这些合并的瓷砖(在我们的示例中为8)将不会在这个倾斜上合并它们自己。你可以在上面的 GIF 动画中找到上面列出的三条规则的应用程序，所以请仔细看几遍，以便更好地理解这些规则。

1. Two tiles of the same value *merge* into one tile containing double the initial number.
2. A tile that is the result of a merge will not merge again on that tilt. For example, if we have [X, 2, 2, 4], where X represents an empty space, and we move the tiles to the left, we should end up with [4, 4, X, X], not [8, X, X, X]. This is because the leftmost 4 was already part of a merge so should not merge again.
3. When three adjacent tiles in the direction of motion have the same number, then the leading two tiles in the direction of motion merge, and the trailing tile does not. For example, if we have [X, 2, 2, 2] and move tiles left, we should end up with [4, 2, X, X] not [2, 4, X, X].

如果倾斜没有改变棋盘状态，那么就不会随机生成新的棋盘。否则，一个单一的随机生成的瓷砖将被添加到一个空的正方形的董事会。注意: 您的代码将不会添加任何新的磁贴！这部分我们已经帮你做好了。

你可能还注意到屏幕底部有一个“ Score”字段，每次移动都会更新。分数不会总是改变每一个动作，但只有当两个瓦片合并。你的代码需要更新分数。

每次两块瓷砖合并形成一个更大的瓷砖，玩家在新瓷砖上获得点数。游戏结束时，当前玩家没有可用的移动(没有倾斜可以改变棋盘) ，或移动形成一个正方形包含2048。你的代码将负责检测游戏何时结束。

“最大得分”是用户在游戏中获得的最大得分。它直到游戏结束才会更新，这就是为什么在动画 GIF 例子中它仍然是0的原因。

它这里为了测试你有没有弄懂这个游戏规则，特意弄了一个Google测试表单来检查，链接在下面。

[2048 Basic Mechanics Quiz (google.com)](https://docs.google.com/forms/d/e/1FAIpQLSeqyhGv2Fpa6HtUfWV4iR71f7pGW6TmRmvtH-X0FXq1KfvE7A/viewform?usp=send_form)

### 简单理解版

+ 收指向一个方向滑动，所有格子都会向那个方向运动。
+ 相同数字的两个格子，碰撞在一起数字会相加，比如2个5碰撞在一起就会产生10.
+ 每次滑动的时候，空白处会随机刷新一个数字的格子。
+ 当界面不可运动时（当界面全部被数字填满时），游戏结束；当界面中最大数字是2048时，游戏胜利。

## Assignment Philosophy and Program Design

在这个项目中，我们将向您提供大量的入门代码，它们使用了许多我们还没有涉及到的 Java 语法，甚至还有一些我们在类中永远不会涉及到的语法。

这里的想法是，在现实世界中，你经常需要使用你不完全理解的代码库，并且需要做一些修改和实验来得到你想要的结果。别担心，当我们下周开始项目1时，你将有机会从头开始。

下面，我们将描述给定框架代码的体系结构背后的一些想法，这些框架代码是由 Paul Hilfinger 创建的。了解每一个细节并不重要，但你可能会发现它很有趣。

骨架展示了两种常用的设计模式: 模型-视图-控制器模式(MVC)和观察者模式。

MVC 模式将我们的问题分为三个部分:

该模型代表所代表的主题事项并对其采取行动——在这种情况下，包括棋盘游戏的状态和可以对其进行修改的规则。我们的模型驻留在 Model、 Side、 Board 和 Tile 类中。模型的实例变量完全决定了游戏的状态。注意: 您只需要修改 Model 类。

模型的视图，该视图向用户显示游戏状态。我们的视图驻留在 GUI 和 BoardWidget 类中。

游戏的控制器，它把用户的动作转化为模型上的操作。我们的控制器主要驻留在 Game 类中，尽管它也使用 GUI 类来读取击键。

MVC 模式不是61B 的主题，在考试或未来的项目中，你也不需要知道或理解这种设计模式。

利用的第二种模式是“观察者模式”。基本上，这意味着模型实际上不会向视图报告更改。相反，视图将自己注册为模型对象的观察者。这是一个有点高级的话题，所以我们在这里不提供额外的信息。

现在我们来看一下你们将要交互的不同类。

### Tile

这个类表示板上的带编号的瓷砖。如果一个 Tile 类型的变量为 null，那么它将被视为板上的一个空的 Tile。您不需要创建这些对象中的任何一个，但是您需要了解它们，因为您将在 Model 类中使用它们。这个类中您需要使用的唯一方法是。Value () ，它返回给定块的值。例如，如果 Tile t 对应于值为8的 Tile，那么 t.value ()将返回8。

### Side

这个是一个枚举类，里面定义了四条边：

+ NORTH
+ SOUTH
+ EAST
+ WEST

### Model

这个类表示game的全部状态，一个对象表示一个2048游戏，它具有状态的实例变量，比如u松油对象在哪，得分是多少等等以及各种方法，当进入这个项目的第四个最终任务(编写方法)时，其中一个挑战是找出这些方法和实例变量中哪些是有用的。

### Board

这个类表示tiles的board本身，它有三个方法，你可以使用setViewingPerspective，tile，getRandomNonNullTile，在这次作业中，只需要编写Model.java就可以了，

## Getting Started 准备工作

在做project0之前，需要先把Lab1搞定，Lab1就是配置IDE和git的，难度不大，但是不晓得student为什么没有开放，所以clone下面的链接就好了。

```
https://github.com/Berkeley-CS61B/skeleton-sp21
```

打开project0，配置下JDK，这里官方说的很清楚了，我用的是Windows。

![ProjectStructure](//CS61BNote/images/ProjectStructure.PNG)

SP21里面的JDK版本是15，我自己的是11，目前没有啥问题，如果配置好的话，Main文件是可以运行的，并且会出现下面这个图片：

![game](//CS61BNote/images/game.PNG)

接下来就要开始动手写代码了。

## Your assignment 

我们要在这个项目中修改和完成类，其他的都已经弄好，官方给出的建议顺序是，先做前两个，因为前两个比较简单，后面两个难度比较高，需要3到10个小时完成。

### public static boolean emptySpaceExists

这个方法在Model类中。

```java
public static boolean emptySpaceExists(Board b) {
    // TODO: Fill in this function.
    return false;
}
```

注释和文档里面说的很清楚了，如果给定板中的任何一个块为空，则此方法应返回 true，空格存储为空，可以直接暴力枚举板中的每一个位置，来判断是否为空，如果为空则直接返回true，不为空则返回false。

还没完成这个方法的时候，测试用例只能过2个，一共8个，如果这个方法是对的时候，则8个都能通过，对应的测试类TestEmptySpace

![test](D:\CS61BNote\images\test.PNG)

### public static boolean maxTileExists

在棋盘中的任何一个棋子如果都等于获胜棋子的值2048，那这个方法应该返回2048，建议不要直接使用2048这个数字，而是使用MAX _ piece这个已经定义好的值。

因为像2048这样直接使用的值属于魔术数字，简单来说会导致可读性下降，关于魔术数字的定义如下(引用维基百科):

```
在程序设计中，魔术数字（magic number）可能指：
缺乏解释或命名的独特数值。常常在程序中出现多次，并且可以（从规范上而言也应当）被有名字的常量取代。
用于识别一个文件格式或协议类型的一段常量或字符串，例如UNIX的特征签名。
不易于其他值混淆的值，例如UUID。
```

对应的测试类:TestMaxTileExists

![test2](D:\CS61BNote\images\test2.PNG)

这里的思路很简单，首先2048的规则是这样的，当棋盘上没有空的位置的时候，并且最大值是2048，游戏就会获得胜利，根据这个规则，先遍历一下棋盘，在判断一下值是不是等于2048就可以了。

### public static boolean atLeastOneMoveExists(Board b)

这个方法(任务)更具有挑战性，它的任意移动是有效的，那它返回true，任意有效的移动是指：假设有一个按键，可以上下左右移动，当用户在玩2048的时候，它可以至少移动一个位置，那这样的按键就是称为有效移动。

用人话来讲就是，只要还能移动，那这个移动就是有效移动，如果不能在移动了，棋盘已经满了，那就不能在移动了，就是无效移动。

对应的测试类是TestAtLeastOneMoveExists。

首先，棋盘为空的时候，不管怎么移动，都是有效的。

当棋盘满了，并且数字为2048的时候，就不能在移动了。

当棋盘还有位置的时候，则开始移动，如果两个值相等就相加，接着重复以上几个步骤。

![test3](D:\CS61BNote\images\test3.PNG)

对应测试类: TestAtLeastOneMoveExists

### Main Task: Building the Game Logic

Boss关，完成这一关，2048这款游戏也就可以玩了，在 public boolean tilt(Side side)中编写代码。

