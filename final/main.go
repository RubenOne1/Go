package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<style>
    html, body, h1 {
        padding: 0;
        border: 0;
        margin: 0;
        box-sizing: border-box;
    }
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        background-image: url("https://images.designtrends.com/wp-content/uploads/2016/06/08052216/Rain-with-Dark-Background.jpg");
        height: 100vh;
    }
    div{
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      background-image: linear-gradient(indigo, darkslateblue, blue);
      height: 70vh;
    }
    h1 {
        font-size: 8rem;
        color: white;
    }
    span{
      color: slategray;
      font-size: 1rem;
    }
</style>
<body>
<main>
  <h1>Remedy</h1>

</main>
<div>
<h1>Emblem</h1>
<span>
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
  This is just sound. This is just thought. This is just noise. This is just me.
</span>
</div>
</body>
</html>
            `)
}
