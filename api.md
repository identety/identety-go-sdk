# App

Methods:

- <code title="get /">client.App.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#AppService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Clients

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#Client">Client</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientListResponse">ClientListResponse</a>

Methods:

- <code title="post /clients">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientNewParams">ClientNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientUpdateParams">ClientUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clients">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientListParams">ClientListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientListResponse">ClientListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#ClientService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#User">User</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserListResponse">UserListResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserListParams">UserListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go">identety</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orgs

Methods:

- <code title="get /org/{id}">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Roles

Methods:

- <code title="get /role/{id}">client.Roles.<a href="https://pkg.go.dev/github.com/stainless-sdks/identety-go#RoleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
