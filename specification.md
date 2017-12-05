# メモ

## 基本
- /etc/resolv.conf にDNSサーバーが記載されている
	- macだと下記URLの方法で変更可能
		- https://pc-karuma.net/mac-network-dns-server/
- 53番portで待ち受けしている
- resolv.conf記載 + 53番待受で見に来てくれないか調査
	- dig?
- ブラウザ/curl等の名前解決のルートを調査する
- hostsファイル
	- ホスト名とIPのマッピング(DNS?)
	- 本質的には同じ
	- DNSは大規模になった場合に利用するイメージ
