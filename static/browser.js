const scroller = document.querySelector("ul")
let topImg = 0
let bottomImg = scroller.childElementCount - 1

/**
 * 
 * @param {string} path 
 * @param {boolean} start 
 */
const addImage = (path, start=false) => {
    const li = document.createElement("li")
    const img = document.createElement("img")
    img.src = "/image/" + path
    img.alt = path
    img.loading = "lazy"
    li.appendChild(img)
    if (start) {
        scroller.prepend(li)
    } else {
        scroller.append(li)
    }
    console.log(scroller.childElementCount)
}


/**
 * 
 * @param {IntersectionObserverEntry} entry 
 */
const handleOne = async (entry) => {
    Observer.unobserve(entry.target)
    if (entry.isIntersecting) {
        const nextImg = await getImageURL(bottomImg + 1)
        if (nextImg) {
            addImage(nextImg)
            bottomImg++
        }
    } else if (scroller.childElementCount > 20) {
        entry.target.remove()
        bottomImg--
    }
    Observer.observe(scroller.lastElementChild)
}

/**
 * @type {IntersectionObserverCallback}
 */
const handle = async (entries) => {
    handleOne(entries[0])
}

const Observer = new IntersectionObserver(handle, {
    root: scroller,
    rootMargin: '100px',
    threshold: 0
})

/**
 * 
 * @param {number} num 
 * @returns {[string]}
 */
const getImageURL = async (num) => {
    const res = await fetch("/list?limit=1&offset="+num)
    const json = await res.json()
    return json[0] 
}

Observer.observe(scroller.lastElementChild)

/**
 * 
 * @param {IntersectionObserverEntry} entry 
 */
const handleTopOne = async (entry) => {
    topObserver.unobserve(entry.target)
    if (entry.isIntersecting && topImg > 0) {
        const nextImg = await getImageURL(topImg - 1)
        if (nextImg) {
            addImage(nextImg, start=true)
            topImg--
        }
    } else if (scroller.childElementCount > 20){
        entry.target.remove()
        topImg++
    }
    topObserver.observe(scroller.firstElementChild)
}

/**
 * @type {IntersectionObserverCallback}
 */
const handleTop = (entries) => {
    entries.forEach(handleTopOne)
}

const topObserver = new IntersectionObserver(handleTop, {
    root: scroller,
    rootMargin: '200px',
    threshold: 0
})