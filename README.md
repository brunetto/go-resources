# Go Resources

## Intro

* [Steve Fracia - The Legacy of Go](https://www.youtube.com/watch?v=k0VsfTAqqEA&t=3s&ab_channel=GoLabconference)
* [Rob Pike - Simplicity is Complicated](https://www.youtube.com/watch?v=k9Zbuuo51go&ab_channel=CodingTech)
* [The Zen of Go](https://dave.cheney.net/2020/02/23/the-zen-of-go)
* [Less is more](https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)
* [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)
* [Go modules & supply chain](https://go.dev/blog/supply-chain)
* [Venice Gophers slides](https://docs.google.com/presentation/d/1FUZhLV9QC5EZBI4yUm3UCiyXChgkWyv0CtBZ6oTZM_I/edit?usp=sharing)
* [JustForFunc](https://www.youtube.com/c/JustForFunc)
* [Go 101](https://go101.org/)
* [Learning resources](https://ryan0x44.medium.com/resources-for-learning-go-79fd2de0d6ae)
* [Official Go Tour](https://go.dev/tour/welcome/1)
* [Case studies, official site](https://go.dev/solutions/#case-studies)
    * [Zalando Skipper](https://github.com/zalando/skipper)
* [Compatibility promise](https://go.dev/doc/go1compat)
* [Patterns](https://blogtitle.github.io/some-useful-patterns/)
* [awsome Go (resources collection)](https://github.com/avelino/awesome-go)
* [Go experience reports](https://github.com/golang/go/wiki/ExperienceReports)
* [Go style guide](https://go.dev/doc/effective_go)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* [Tests Best Practices](https://github.com/golang/go/wiki/TestComments)
* [Language specs](https://go.dev/ref/spec)

## Packages and project structure

* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
* [Style guideline for Go packages](https://rakyll.org/style-packages/)
* [Package names](https://go.dev/blog/package-names)
* [Ashley McNamara + Brian Ketelsen. Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0&ab_channel=GopherConRussia)
* [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0&ab_channel=GopherAcademy) 
* [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* [Packages as layers](https://www.gobeyond.dev/packages-as-layers/)
* [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg&ab_channel=GopherConEurope)
* [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

## Web

* [Bill Kennedy - You Want To Build a Web Service?](https://www.youtube.com/watch?v=IV0wrVb31Pg&ab_channel=GoLabconference)
   * [Ardan Labs repo](https://github.com/ardanlabs/service#more-about-go)
* [exposing Go on the internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)
* [1M RPM with Go](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)
* [REST Servers in Go](https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/)

## Slices 

* [Slices Gotchas](https://blogtitle.github.io/go-slices-gotchas/)
* [Slices from the ground up](https://dave.cheney.net/2018/07/12/slices-from-the-ground-up)
* [Slice utils](https://pkg.go.dev/golang.org/x/exp/slices)

##  Concurrency

* [Concurrency patterns 1](https://blogtitle.github.io/go-advanced-concurrency-patterns-part-1/)
* [Concurrency patterns 2](https://blogtitle.github.io/go-advanced-concurrency-patterns-part-2-timers/)
* [Concurrency patterns 3](https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/)

## Generics 

* https://go.dev/blog/when-generics
* [Generics intro and opinions](https://blogtitle.github.io/go-generics/)
* [Tutorial](https://go.dev/doc/tutorial/generics)
* [Generics Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
* https://go101.org/generics/101.html#index
* [Generics can make your Go code slower](https://planetscale.com/blog/generics-can-make-your-go-code-slower)
* https://github.com/bradenaw/juniper
* https://eli.thegreenplace.net/2022/faster-sorting-with-go-generics/
* https://github.com/golang/proposal/blob/master/design/generics-implementation-dictionaries-go1.18.md
* https://encore.dev/blog/go-1.18-generic-identifiers
* https://christine.website/blog/gonads-2022-04-24
* https://medium.com/@jon_43067/go-generics-and-concurrency-d0dccab73a73

## Errors

* [Errors are values](https://go.dev/blog/errors-are-values)
* [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
    * [Video](https://www.youtube.com/watch?v=lsBF58Q-DnY&ab_channel=GopherAcademy)
* [Idiomatic Go Tricks](https://www.youtube.com/watch?v=yeetIgNeIkc&t=257s&ab_channel=GopherConUK)
* [Why Go gets exceptions right](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)
* [Dismissed try proposal](https://github.com/golang/go/issues/32437)
* [An Open Letter To The Go Team About Try](https://www.ardanlabs.com/blog/2019/07/an-open-letter-to-the-go-team-about-try.html)
*   [Ref tweet](https://mobile.twitter.com/goinggodotnet/status/1144628423199002624)
* [proposal: leave "if err != nil" alone?](https://github.com/golang/go/issues/32825)
* [Error Handling — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md)
* [Considerations on error handling](https://blogtitle.github.io/considerations-on-error-handling/)
* https://levelup.gitconnected.com/better-error-handling-in-golang-theory-and-practical-tips-758b90d3f6b4
* https://earthly.dev/blog/golang-errors/

## JSON

* https://pkg.go.dev/github.com/Jeffail/gabs#section-readme

## Auth/authz

* [Go HTTPS servers with TLS](https://eli.thegreenplace.net/2021/go-https-servers-with-tls/)

## Logging

* https://github.com/rs/zerolog
* https://betterstack.com/community/guides/logging/zerolog/

## Vulnerabilities 

* https://go.dev/blog/vuln

## Context

* [Context Package Semantics In Go](https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html)
* [Getting Started with Go Context](https://dev.to/gopher/getting-started-with-go-context-l7g)
* [Context is for cancelation](https://dave.cheney.net/2017/01/26/context-is-for-cancelation)
* [Context isn’t for cancellation](https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation)
* [Go Concurrency Patterns: Context](https://go.dev/blog/context)
* [Cancelation, Context, and Plumbing](https://talks.golang.org/2014/gotham-context.slide#18)
<!--* [Understanding the context package in golang](http://p.agnihotry.com/post/understanding_the_context_package_in_golang/)
* [It's time to understand Golang Contexts.](https://www.linkedin.com/pulse/its-time-understand-golang-contexts-lucas-schenkel-schieferdecker/)
* [Practical Go Lessons: context](https://www.practical-go-lessons.com/chap-37-context)
* [Golang and context - an explanation](https://pauldigian.com/golang-and-context-an-explanation)
* [Go Context 101](https://medium.com/codex/go-context-101-ebfaf655fa95)
* [Context in Golang!](https://levelup.gitconnected.com/context-in-golang-98908f042a57)
* [How to correctly use context.Context in Go 1.7](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39) -->

## IDs

* [How we used Go 1.18 when designing our Identifiers](https://encore.dev/blog/go-1.18-generic-identifiers)
* [Globally Unique ID Generator](https://github.com/rs/xid) (this seems to be the best)
* [UUID](https://github.com/gofrs/uuid)
* [ksuid](https://github.com/segmentio/ksuid)

## Fuzzing 

* [Docs](https://go.dev/doc/fuzz/)
* [Tutorial](https://go.dev/doc/tutorial/fuzz)

## Misc

* [Code Review Developer Guide](https://google.github.io/eng-practices/review/) 
* [A Lodash-style Go library](https://github.com/samber/lo)
* [Go Does Not Need a Java Style GC](https://itnext.io/go-does-not-need-a-java-style-gc-ac99b8d26c60)
* [Go: A Documentary](https://golang.design/history/)
* [Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)
* [GothamGo 2018 - Things in Go I Never Use by Mat Ryer](https://www.youtube.com/watch?v=5DVV36uqQ4E&ab_channel=NautHumanProductions)
* [Code: Align the happy path to the left edge](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88)
* [Awsome Go: A curated list of awesome Go frameworks, libraries and software](https://github.com/avelino/awesome-go)

# Interesting libraries and projects

* [Build cross-platform GUI applications using Go + HTML + CSS + JS](https://wails.io/)
   * https://wails.io/blog/wails-v2-released/
* [Commit Message Emoji 👋](https://github.com/dannyfritz/commit-message-emoji)
* Image manipulation
   * [Caire: Content aware image resize library](https://github.com/esimov/caire)
   * [Imagor](https://github.com/cshum/imagor) 
   * [bild](https://github.com/anthonynsimon/bild)
   * [bimg](https://github.com/h2non/bimg)
* [LazyGit](https://github.com/jesseduffield/lazygit)
* [Got.](https://github.com/melbahja/got) (~ curl and wget)
* [An alternative JSON decoder for Go.](https://github.com/pkg/json)
* [The Go+ language for engineering, STEM education, and data science](https://github.com/goplus/gop)
* [Thist](https://github.com/bsipos/thist) (terminal histograms)
* https://github.com/shivamMg/trie
* https://taskfile.dev/
* https://github.com/go-playground/validator/
