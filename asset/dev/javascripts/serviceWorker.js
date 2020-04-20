const staticFun = "fun-v1";
const assets = [
    "/",
    "/stylesheets/styles.css",
    "/javascripts/main.js",

    // Fonts
    "/fonts/fa-brands-400.eot",
    "/fonts/fa-brands-400.ttf",
    "/fonts/fa-brands-400.woff2",
    "/fonts/fa-regular-400.svg",
    "/fonts/fa-regular-400.woff",
    "/fonts/fa-solid-900.eot",
    "/fonts/fa-solid-900.ttf",
    "/fonts/fa-solid-900.woff2",
    "/fonts/fa-brands-400.svg",
    "/fonts/fa-brands-400.woff",
    "/fonts/fa-regular-400.eot",
    "/fonts/fa-regular-400.ttf",
    "/fonts/fa-regular-400.woff2",
    "/fonts/fa-solid-900.svg",
    "/fonts/fa-solid-900.woff",

    // Images
    "/images/icons/icon-72x72.png",
    "/images/icons/icon-96x96.png",
    "/images/icons/icon-128x128.png",
    "/images/icons/icon-144x144.png",
    "/images/icons/icon-152x152.png",
    "/images/icons/icon-192x192.png",
    "/images/icons/icon-384x384.png",
    "/images/icons/icon-512x512.png",
];

self.addEventListener("install", installEvent => {
    installEvent.waitUntil(
        caches.open(staticFun).then(cache => {
            cache.addAll(assets);
        })
    );
});

self.addEventListener("fetch", fetchEvent => {
    fetchEvent.respondWith(
        caches.match(fetchEvent.request).then(res => {
            return res || fetch(fetchEvent.request)
        })
    );
});