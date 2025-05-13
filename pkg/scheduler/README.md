# MEMO

## TODO

### This repo

- [x] [KWOK](https://kwok.sigs.k8s.io/)の使い方を調べる
- [x] kwokでカスタムスケジューラを動かす方法を調査
- [ ] スクラッチでkube-schedulerを実装
  - [ ] [自作して学ぶKubernetes Scheduler](https://engineering.mercari.com/blog/entry/20211220-create-your-kube-scheduler/)を読む
  - [ ] eventHandlerとactiveQを実装する
    - [ ] KWOKとログでactiveQにPodが追加されることを確認できないか
  - [ ] unschedulableQを実装する
  - [ ] backoffQを実装する
  - [ ] その他を実装する
    - [ ] inFlightEvents, inFlightPodsの対応
    - [ ] logger, metricsの対応
  - [ ] percentageOfNodesToScoreの仕組みを実装する
  - [ ] power of twoを実装してみる
- [ ] ユニットテストを書く＆テストケースを追加してPRを出す
- [ ] ランダムスケジュールをするプラグインを自作してkwok上でkube-schedulerの代わりに動かす
  - [ ] 全てのNodeを取得するPreScoreプラグインを実装する
  - [ ] ランダムにNodeを一つ選ぶScoreプラグインを実装する
  - [ ] Reserveプラグインを実装する
  - [ ] PostBindでログを出力する
  - [ ] スケジューラの実行を確認する
- [ ] デフォルトで使用されているプラグインを自作してkwok上でkube-schedulerの代わりに動かす
  - [ ] [NodeResourcesFit plugin incorrectly computes requested resources #130445](https://github.com/kubernetes/kubernetes/issues/130445)
- [ ] インテグレーションテストを自作してテストを書く＆テストケースを追加してPRを出す
- [ ] scheduler_perfを自作してテストを書く＆テストケースを追加してPRを出す
  - [test: Add test case for createNodesOp #131607](https://github.com/kubernetes/kubernetes/pull/131607)
  - [add tests for scheduler-perf itself #127745](https://github.com/kubernetes/kubernetes/issues/127745)
- [ ] プリエンプションをするプラグインを自作してkwok上でkube-schedulerの代わりに動かす
  - [ ] 非同期プリエンプションも実装する
  - [猫でもわかるPod Preemption](https://speakerdeck.com/ytaka23/kubernetes-meetup-tokyo-10th)
  - [A Deeper Dive of kube-scheduler](https://www.awelm.com/posts/kube-scheduler/)
- [ ] Gang Schedulingをするプラグインを自作してkwok上でkube-schedulerの代わりに動かす
  - coschedulingとpfnetの実装を比較して書く
  - [ ] pfnetの実装を参考に、scheduler-plugins/coschedulingを改善
- [ ] [scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator)を自作
- [ ] 本物のscheduling frameworkを使ってプラグインを作って動かす
- [ ] E2Eテストを動かしてみる
- [ ] これまでの話をまとめてZennとMediumに「自作して学ぶkube-scheduler」シリーズを連載する

### Other

- issueをウォッチして、議論に参加して、このissueは俺がやりたいと宣言すれば実装できる
  - 常に下記のリポジトリをウォッチし、キャッチアップする
- kubernetes
  - kube-scheudler
    - 網羅されているテストケースを列挙し、漏れがないかを分析して、テストを追加する
    - パフォーマンスボトルネックを見つけてコードを改善する
      - プリエンプションやgang scheduling周りなど
      - 処理のシーケンス図を作って、各処理にかかったCPU時間を計測すると見えてくるものがあるのでは
      - プリエンプションを非同期に実行するのを実装してみてパフォーマンス比較
      - parallel-schedulerを実装してみてパフォーマンスを比較してみたら何かわかるかも
      - APMを実装する
        - 結局はどの関数が同期的、非同期的に実行されて、それぞれかかったCPU時間と累計CPU時間、メモリ使用量、他コンポーネントなどとのI/oにかかって時間などがすぐに分かって欲しいわけで、それがわかるビジュアライズツールを作るとか？
        - 縦軸を絶対CPU時間にして、各関数の呼び出しを時系列順に並べて、各関数でどのぐらいのCPU時間を使い、絶対CPU時間はこのぐらい経ってるというのをわかりやすく可視化する（マルチスレッドに対応）
      - APMのデザインドキュメントを作ってissueを立てて提案
    - percentageOfNodesToScoreを用いた際のスキュー問題を解決する
      - [Scheduler is not balancing properly the pods across the nodes in big clusters (>200 nodes) in quick massive scale ups #130692](https://github.com/kubernetes/kubernetes/issues/130692)
      - [Scheduler: use the power of two random choices to select nodes #86630](https://github.com/kubernetes/kubernetes/issues/86630)
      - KubeSchedulerConfigurationのparallelismではNodeに対しての横断的な処理の並列数を指定できる。これが原因で卒論でScheduling Cycleを並列化してもあまり効果がなかった可能性がある。
      - power-of-twoを実装してみて評価してみる？
      - Podの配置の質を評価するテストってある？
        - 例えばbin packingやspread戦略を取る場合に、最適な配置を正解として実際のスケジューリングの結果と比較して、どのぐらいの精度で配置されているかを評価するテストが欲しい
        - またはspreadならノード間の最大の差がどのぐらいになるのか、bin packingならノードの使用率がどれくらいになるのかで評価できる
        - 独自のプラグインを実装した際にちゃんと意図した通りに動作するのかを確かめる必要があるので。そのようなテストを実行できるようにする
        - インテグレーションテストなどでそのようなテストケースはあるのだろうか？
        - [Scheduler is not balancing properly the pods across the nodes in big clusters (>200 nodes) in quick massive scale ups #130692](https://github.com/kubernetes/kubernetes/issues/130692)
  - scheduler_perf
    - gang schedulingのシナリオを追加？
- kube-scheduler-simulator
  - gang schedulingを検証するのに役立つ機能を追加?
  - 複数のスケジューラを同時に動かしてシミュレートすることってできたっけ？できないのであればできるようにするとか？
- scheduler-plugins
- lws
  - gang schedulingの実装
- kwok

## メモ

- client-go：APIサーバーとやり取りをするためのクライアントライブラリ
- Shared Informer Factory：client-goで提供されている型で、K8sのリソース（NodeやPodなど）を効率的に監視（watch）するための仕組み。１つのSharedInformerFactoryインスタンスから様々なリソース用のInformerを作成できる。APIサーバーへのリクエストを減らすために、リソースの状態をローカルキャッシュとして保持する。リソースの変更（追加、更新、削除）時に登録してあるコールバック関数（EventHandler）を呼び出すことができる。sharedとは、複数のInformer（監視・キャッシュ機構）が、同じリソースの監視やキャッシュを共有するという意味で、たとえば、Podの情報を監視するInformerが2つあった場合、個別にInformerを作るとAPIサーバーへのwatchリクエストが2回発生するが、SharedInformerFactoryを使うと1回のwatchリクエストと1つのキャッシュを複数のInformerで共有できる。
- Informer：K8sのリソースを監視するためのコンポーネント。特定のリソース（NodeやPodなど）を監視し、変更があった場合に通知を受け取ることができる。
- EventHandler：[Extend Kubernetes via a shared informer](https://www.cncf.io/blog/2019/10/15/extend-kubernetes-via-a-shared-informer/)

```zsh
cd pkg/scheduler
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
