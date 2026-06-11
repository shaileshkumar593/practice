from flask import Flask
app = Flask(__name__)

@app.route('/hello/<name>')
def hello_world(name):
   return f'Hello World {name}'

@app.route('/blog/<int:postID>')
def show_blog(postID):
   return f'Blog Number is {postID}'

@app.route('/rev/<float:revNo>')
def revision(revNo):
   return f'Revision Number is {revNo}'

if __name__ == '__main__':
   app.run( debug=True)
