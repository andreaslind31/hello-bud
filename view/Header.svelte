<script>
    let src = "https://news.ycombinator.com/y18.gif";
</script>

<div class="header">
    <img class="img-logo" {src} alt="logo" />
    <a href="/"><h1 class="header-text">Hacker news</h1></a>
    <div class="links">
        <a href="/newest">Newest</a>
        <a href="/users">Users</a>
    </div>
</div>

<style>
    .header {
        display: flex;
        align-items: center;
        padding: 10px;
    }
    .img-logo {
        display: block;
        height: 30px;
        width: 30px;
    }
    .header-text {
        margin-left: 10px;
        font-weight: 600;
        font-size: 1rem;
    }
    h1:hover {
        text-decoration: none;
    }
    .links {
        margin-left: 10px;
        display: flex;
        align-items: center;
    }
    .links > a {
        margin-left: 10px;
        font-size: 75%;
    }
    .links > a:hover {
        text-decoration: underline;
    }
    a {
        text-decoration: none;
        color: inherit;
    }
</style>
