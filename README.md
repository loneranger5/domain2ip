# domain2ip
concurrent domain-to-IP resolver written in Go using `net.LookupIP` and `github.com/sourcegraph/conc` for parallel execution. It reads a list of domain names from a file and outputs resolved IPv4 addresses to both the console and a specified output file.
