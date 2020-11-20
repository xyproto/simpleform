# SimpleForm

SimpleForm is a language for constructing HTML forms out of very little text.

The conversion is fast enough to happen on the fly, when serving HTML pages.

Here's a simple login form:

```template
Login

Welcome dear user!

Username: {{ username }}
Password: {{ password }}

[Login](/login)
```

* The first line is recognized as the title, and is used both for the title (like this: `<title>Login</title>` and as a title in the body, like this: `<h1>Login</h1>`.
* The lines with `{{` and `}}` are recognized as single line input fields, where the label is the word before `:`.
* The `[Login](/login)` is recognized as a button that can submit the contents of the form as a POST request to `/login`.

Here's the generated output from the above form file:

```html
<!doctype html>
<html lang="en">
  <head>
    <title>Login</title>
  </head>
  <body>
    <h1>Login</h1>
    <p>Here you can log in. Welcome!</p>
    <form method="POST">
      <label for="username">Username:</label><input type="text" id="username" name="username"><br><br>
      <label for="password">Password:</label><input type="password" id="password" name="password"><br><br>
      <input type="submit" formaction="/login" value="Login"><br><br>
    </form>
  </body>
</html>
```

* If the input ID starts with `password` or `pwd`, then the input type `"password"` is used.
* Multiple buttons can be provided on a single line.
* All text that is not recognized as either the title or as form elements is combined and returned in a `<p>` tag.

## General Info

* Version: 0.1.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
