# Exercise 7.18
Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes. Each element node has a slice of child nodes.
You may find the following declarations helpful.


---
# 練習問題 7.18
トークンに基づくデコーダのAPIを使用して、任意のXMLのドキュメントを読み込んで、そのドキュメントを表す総称的なノードのツリーを構築するプログラムを書きなさい。ノードには二種類あり、`CharData`ノードはテキスト文字列を表し、`Element`ノードは名前付き要素とその属性を表します。それぞれの要素のノードは子ノードのスライスを持ちます。

````go
import "encoding/xml"

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
  Type     xml.Name
  Attr     []xml.Attr
  Children []Node
}
````


# Result

````shel
budougumi0617@~/git/gotraining/ch07/ex18 (remainingwork@GoTraining)
$  cat ../ex17/ex17.xml && echo ----------- && cat ../ex17/ex17.xml| go run decoder.go
<?xml version="1.0" encoding="UTF-8" ?>
<div>
  <div>
    <name>Bad attribute</name>
    <url>???</url>
  </div>
  <div2 id="foo">
    <name>Correct attribute</name>
    <url>OK!!</url>
  </div2>
  <div>
    <name>Bad attribute</name>
    <url>???</url>
  </div>
</div>
-----------
Result:
&main.Element{Type:xml.Name{Space:"", Local:"div"}, Attr:[]xml.Attr{}, Children:[]main.Node{"\n  ", (*main.Element)(0xc8200740f0), "\n  ", (*main.Element)(0xc8200741e0), "\n  ", (*main.Element)(0xc8200742d0), "\n"}}
````
