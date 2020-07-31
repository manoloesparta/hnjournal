import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return CupertinoApp(
      debugShowCheckedModeBanner: false,
      title: 'Hacker News Archive',
      home: MyHomePage(title: 'Hey'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: CupertinoNavigationBar(
        middle: Text('Hacker News Jorunal'),
      ),
      child: Center(
        child: Container(
          margin: EdgeInsets.only(top: 150),
          child: Column(
            children: <Widget>[
              Text('Hola', textScaleFactor: 2),
              Text('Hola', textScaleFactor: 2),
              Text('Hola', textScaleFactor: 2),
            ],
          ),
        )
      )
    );
  }
}