# duckend

1. Create `.env` file containing the following:

- `STRIPE_KEY`
- `PRICE_ID`

2. Build docker image: `docker build -t duckend .`

3. Run docker image: `docker run -d -p 80:80 duckend`
