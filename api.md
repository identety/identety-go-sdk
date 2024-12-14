# Clients

Response Types:

- <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#Client">Client</a>
- <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientListResponse">ClientListResponse</a>

Methods:

- <code title="post /clients">client.Clients.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientNewParams">ClientNewParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientUpdateParams">ClientUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /clients">client.Clients.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientListParams">ClientListParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientListResponse">ClientListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /clients/{id}">client.Clients.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#ClientService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#User">User</a>
- <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserListResponse">UserListResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserListParams">UserListParams</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/{id}">client.Users.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk">identety</a>.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orgs

Methods:

- <code title="get /org/{id}">client.Orgs.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Roles

Methods:

- <code title="get /role/{id}">client.Roles.<a href="https://pkg.go.dev/github.com/identety/identety-go-sdk#RoleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
