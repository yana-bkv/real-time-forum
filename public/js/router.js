const routes = {
    "/": {
        template: "/templates/home.html",
    },
    "/login": {
        template: "/templates/login.html",
    },
    "/register": {
        template: "/templates/register.html",
    },
};

function navigate(event, path) {
    event.preventDefault(); // Prevent page reload
    history.pushState({}, "", path); // Change URL without reloading
    router();
}

const route = (event) => {
    event = event || window.event;
    event.preventDefault();
    window.history.pushState({},"", event.target.href);
    router()
}

const router = async () => {
    const content = document.getElementById("content");
    var path = window.location.pathname;
    if (path.length === 0) {
        path = "/";
    }
    const route = routes[path] || "404";
    const html = await fetch(route.template).then((res) => res.text());
    content.innerHTML = html;
}

window.addEventListener("popstate", router);
window.route = route;
window.addEventListener("load", router);

