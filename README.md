# shorturl

Step-1: `git clone git@github.com:bvighnesha/shorturl.git`

Step-2: `cd shorturl`

Step-3: `docker build -t shorturl .`

Step-4: `docker run -p "3000:3000" shorturl`

Step-4: open browser enter `http://localhost:3000/url/{url}`

Example: open browser enter `http://localhost:3000/url/abc.com` will generate `short.url/1R9doEiEshDAGoY`

Short url will be generated and automatically will be redirected to short url.