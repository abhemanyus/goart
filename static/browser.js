const scroller = document.querySelector("ul")
let topImg = 0
let bottomImg = scroller.childElementCount - 1

/**
 * 
 * @param {string} path 
 */
const addImage = (path) => {
    const li = document.createElement("li")
    const img = document.createElement("img")
    img.src = "/image/" + path
    img.alt = path
    li.appendChild(img)
    scroller.appendChild(li)
    if (scroller.childElementCount > 20) {
        while (scroller.childElementCount > 10) {
            scroller.removeChild(scroller.firstChild) 
        }
    }
    console.log(scroller.childElementCount)
}


/**
 * 
 * @param {IntersectionObserverEntry} entry 
 */
const handleOne = async (entry) => {
    if (entry.isIntersecting) {
        const nextImg = await getImageURL(bottomImg + 1)
        if (nextImg) {
            Observer.unobserve(entry.target)
            addImage(nextImg)
            bottomImg++
            topImg++
            setTimeout(() => Observer.observe(scroller.lastElementChild), 500)  
        }
    }
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
    threshold: 0.5
})

/**
 * 
 * @param {number} num 
 * @returns {string}
 */
const getImageURL = async (num) => {
    const res = await fetch("/list?limit=1&offset="+num)
    const json = await res.json()
    return json[0] 
}

Observer.observe(scroller.lastElementChild)