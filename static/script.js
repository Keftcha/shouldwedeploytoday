var main = () => {
    const app = document.getElementById("app")
    app.innerHTML = "Loading..."

    // Fetching data to the api
    fetch("/api/" + (new Date()).getDay())
        .then(resp => {
            resp.json()
                .then(data => displayDailyInfo(data, app))
        })
        .catch(resp => {
            app.innerHTML = "Failed to fetch info."
        })
}

function displayDailyInfo(data, node) {
    node.innerHTML = ""
    var center = document.createElement("center")
    node.appendChild(center)

    var title = document.createElement("h1")
    title.innerHTML = data.title
    title.classList.add("title")
    center.appendChild(title)

    var img = document.createElement("img")
    img.setAttribute("src", data.image)
    img.classList.add("image")
    center.appendChild(img)

    center.appendChild(document.createElement("br"))
    center.appendChild(document.createElement("br"))

    var text = document.createElement("span")
    text.innerHTML = data.text
    text.classList.add("text")
    center.appendChild(text)
}

window.addEventListener("load", main)
