# Standard Library/Golang: API Basic Role-Based Access Control (RBAC) Code Sample

This Golang code sample demonstrates **how to implement authorization** using Auth0 by Okta and the [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) package.

This code sample is part of the ["Auth0 Developer Resources"](https://developer.auth0.com/resources), a place where you can explore the authentication and authorization features of the Auth0 Identity Platform.

The API server is built with the Golang Standard Library and uses the [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) package. This code sample shows you how to accomplish the following tasks:

- Register a Golang API in the Auth0 Dashboard.
- Use Golang middleware to enforce API security policies.
- Perform access control in Golang using a token-based authorization strategy powered by JSON Web Tokens (JWTs).
- Validate access tokens in JSON Web Token (JWT) format using Golang middleware.
- Make authenticated requests to a secure Golang API server.

## Golang Code Sample Specs

This code sample uses the following main tooling versions:

- Golang `v1.21`
- [GO JWT Middleware Auth0 SDK `v2.2.1`](https://github.com/auth0/go-jwt-middleware)

## Quick Auth0 Set Up

**First and foremost, if you haven't already, [sign up for an Auth0 account](https://auth0.com/signup) to connect your API with the Auth0 Identity Platform.**

Next, you need to create an API registration in the Auth0 Dashboard. You'll get two configuration values, the **Auth0 Audience** and the **Auth0 Domain**, that will help connect your API server with Auth0.

### Get the Auth0 audience

- Open the [APIs](https://manage.auth0.com/#/apis) section of the Auth0 Dashboard.

- Click on the **Create API** button and fill out the **"New API**" form with the following values:

```bash
# Name
Hello World API Server
# Identifier
https://hello-world.example.com
```

- Click on the **Create** button.

> Visit the ["Register APIs"](https://auth0.com/docs/get-started/auth0-overview/set-up-apis) document for more details.

When setting up APIs, we also refer to the API identifier as the Audience value. You will use that value in the next sections.

### Get the Auth0 domain

Now, follow these steps to get the **Auth0 Domain** value.

- Open the [Auth0 Domain Settings](https://manage.auth0.com/#/tenant/custom_domains)

- Locate the bold text in the page description that follows this pattern: `tenant-name.region.auth0.com`. That's your Auth0 domain!

- You will use the Auth0 domain in the next section to set up your API server.

> The `region` subdomain (`au`, `us`, or `eu`) is optional. Some Auth0 Domains don't have it.

## Set Up and Run the Golang Project

Start by cloning the Golang project:

```bash
git clone https://github.com/byron-okta/api_standard-library_golang_hello-world.git
```

Make the project directory your current working directory:

```bash
cd api_standard-library_golang_hello-world
```

The `main` branch holds all the code related to implementing token-based authorization to protect resources in a Golang API.

Install the Golang project dependencies:

```bash
go mod download
```

Now, create a `.env` file under the root project directory:

```bash
touch .env
```

Populate it with the following environment variables:

```bash
PORT=6060
CLIENT_ORIGIN_URL=http://localhost:4040
AUTH0_AUDIENCE=YOUR_AUTH0_AUDIENCE
AUTH0_DOMAIN=YOUR_AUTH0_DOMAIN
```

Execute the following command to run the Golang API server:

```bash
go run ./cmd/api-server/
```

## Request API Resources from a Client Application

Let's simulate an essential feature of an API: serving data to client applications.

You can pair this API server with a client application that matches the technology stack that you use at work. Any **"Hello World" client application** can communicate with this **"Hello World" API server sample**.

For example you can pair this API with a Next.js Application.

We got you covered with code samples for Next.js app using App Router and Pages Router, visit the following resources to see how to set the client application using Next.js:

- [Next.js Authentication By Example: Using App Router](https://developer.auth0.com/resources/guides/web-app/nextjs/basic-authentication)
- [Next.js/TypeScript Pages Router Code Sample: Basic Authentication](https://developer.auth0.com/resources/code-samples/web-app/nextjs/basic-authentication/typescript-pages-router)

Each client application sample gives you clear instructions to get it up and running quickly.

Once you set up the client application, log in and visit the protected **"Profile"** page ([`http://localhost:4040/profile`](http://localhost:4040/profile)) to see all the user profile information that Auth0 securely shares with your application using [ID tokens](https://auth0.com/docs/security/tokens/id-tokens).

Then, visit the **"Protected"** page ([`http://localhost:4040/protected`](http://localhost:4040/protected)) or the **"Admin"** page ([`http://localhost:4040/admin`](http://localhost:4040/admin)) to practice requesting protected resources from an external API server using [access tokens](https://auth0.com/docs/security/tokens/access-tokens).

> Be sure that you're running the `basic-authentication-with-api-integration` branch in the Next.js code sample, this branch have the setup to communicate with the API server.

For more information on how to integrate it with an API, please visit the following sections:

- [Integrate Next.js with an API Server](https://developer.auth0.com/resources/guides/web-app/nextjs/basic-authentication#integrate-next-js-with-an-api-server) for Next.js with App Router.
- [Connect the Next.js Code Sample with an API Server](https://developer.auth0.com/resources/code-samples/web-app/nextjs/basic-authentication/typescript-pages-router#connect-the-next-js-code-sample-with-an-api-server) for Next.js with Pages Router
