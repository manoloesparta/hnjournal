import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:http/http.dart' as http;

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return CupertinoApp(
      debugShowCheckedModeBanner: false,
      title: 'Hacker News Archive',
      theme: CupertinoThemeData(
        brightness: Brightness.light
      ),
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key}) : super(key: key);
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {

  List<HackerNew> hns = new List<HackerNew>();

  void _getNew() async {
    hns.clear();

    for(int i  = 0; i < 8; i++) {
      var ip = 'http://64.227.54.150:8080/random';

      var response = await http.get(ip);
      var resjson = json.decode(response.body);
      
      setState(() {
        hns.add(HackerNew(title: resjson['title'], url: resjson['URL']));
      });
    }
  }
  
  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: CupertinoNavigationBar(
        middle: Text('Hacker News Jorunal'),
      ),
      child: Container(
        child: ListView(
          children: <Widget>[
            CupertinoButton(
              onPressed: _getNew,
              child: Text('Load posts'),
            ),
            ...hns,
          ],
        ),
      )
    );
  }
}

class HackerNew extends StatelessWidget {

  final String url;
  final String title;
  HackerNew({Key key, this.title, this.url}) : super(key: key);
  
  _launchURL() async {
    if (await canLaunch(this.url)) {
      launch(this.url);
    } else {
      throw 'Could not launch ${this.url}';
    }
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: _launchURL,
      child: Container(
        padding: EdgeInsets.all(10),
          child: Text(
            this.title,
            textAlign: TextAlign.center,
          ),
      )
    );
  }
}
