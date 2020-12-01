// Vlang Curl CLI Tool | Faster version of curl
// SmallDoink#0666
import os
import net.http

// Banner function to print the help option
fn banner() {
	println("vcurl - vlang port of curl")
	println("-h, --help\t show this help message")
	println('-s, --silent\t silence output')
	println('-v, --verbose\t verbose output')
	println('-M, --method\t <GET\\HEAD\\POST')
	println('-H, --headers\t <headers>')
}

fn main() {
	mut url := ""
	mut ua := "vcurl/1.0"
	mut verbose := false
	mut silent := false
	mut method := http.method_from_str('GET')
	mut headers := ['Accept']
	mut data := ['*/*']
	if os.args.len == 1 {
		banner()
		exit(0)
	}
	for i, arg in os.args {
		if 'http' in arg {
			url = arg
		} else if '-h' in arg || '--help' in arg {
			banner()
			exit(0)
		} else if '-v' in arg || '--verbose' in arg {
			verbose = true
		} else if '-s' in arg || '--silent' in arg {
			silent = true
		} else if '-M' in arg || '--method' in arg {
			method = http.method_from_str(os.args[i+1])
		} else if '-H' in arg || '--headers' in arg {
			headers << os.args[i+1].split(':')[0]
			data << os.args[i+1].split(':')[1].split(' ')[1]
		}
	}
	if url == "" {
		banner()
		exit(0)
	}
	mut req := http.new_request(method, url, '')?
	for i, _ in headers {
		req.add_header(headers[i], data[i])
	}
	req.add_header('User-Agent', ua)
	resp := req.do()?
	if verbose {
		path := req.url.split('/')[3]
		println('> $req.method /$path')
		println('< $resp.status_code')
		for key, val in req.headers {
			println('> $key: $val')
		}
		for key, val in resp.headers {
			println('< $key: $val')
		}
	}
	if !silent {
		println(resp.text)
	}
}
