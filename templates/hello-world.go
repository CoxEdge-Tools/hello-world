package templates

const (
	reqInfoID = "reqInfo"

	webHead = `<html>
  <head>
    <title>CoxEdge</title>
    <style>
      body {
        background-color: white;
        text-align: center;
        padding: 50px;
        font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
      }
      button {
          background-color: #0075a8; 
          border: none;
          color: white;
          padding: 15px 32px;
          text-align: center;
          text-decoration: none;
          display: inline-block;
          font-size: 16px;
      }

      #logo {
        margin-bottom: 40px;
      }
    </style>
  </head>
  <body>
    <h2><a style="text-decoration: none; href=" https:="https://www.coxedge.com/">CoxEdge</a></h2>
    <img id="logo" src="img/logo.svg" alt="logo" width=400 />
    <h1>Hello world!</h1>
    <h3>My hostname is {{.Hostname}}</h3>
    <br>
    <h4>Git Commit</h4>
    <h4>{{.GitCommit}}</h4>`

	webServices = `{{- $length := len .Services }} 
  {{- if gt $length 0 }}
    <div id='Services'>
      <h3>k8s services found {{$length}}</h3>
    {{ range $k,$v := .Services }}
      <b>{{ $k }}</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />
  {{ end }}`

	webDetails = `    <button class='button' onclick='myFunction()'>Show request details</button>
    <div id="` + reqInfoID + `" style='display:none'>
      <h3>Request info</h3>
      <b>Host:</b> {{.Host}} <br />
      <b>Pod:</b> {{.Hostname}} </b><br />
    {{ range $k,$v := .Headers }}
      <b>{{ $k }}:</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />`

	webLinks = `    <div id='Links' class="row social">
      <a class="p-a-xs" href="https://github.com/coxedge"><img src="img/icon-github.svg" alt="github" height="25" width="25"></a>
      <a class="p-a-xs" href="https://twitter.com/CoxComm"><img src="img/icon-twitter.svg" alt="twitter" height="25" width="25"></a>
      <a class="p-a-xs" href="https://www.linkedin.com/showcase/cox-edge><img src="img/icon-linkedin.svg" height="25" alt="linkedin" width="25"></a>
    </div>
    <br />`

	webTail = `    <script>
      function myFunction() {
          var x = document.getElementById("` + reqInfoID + `");
          if (x.style.display === "none") {
              x.style.display = "block";
          } else {
              x.style.display = "none";
          }
      }
    </script>
  </body>
</html>`

	HelloWorldTemplate = webHead + `
` + webServices + `
` + webLinks + `
` + webDetails + `
` + webTail
)
