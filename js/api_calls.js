document.addEventListener('DOMContentLoaded', function () {
    console.log("loading")

    let images = document.querySelectorAll('img');
    let loadedImages = 0;

    function checkAllImagesLoaded() {
        loadedImages++;
        console.log("IMAGES LENGTH:", images.length)
        console.log("LOADED: ", loadedImages)
        if (loadedImages === images.length) {
            // All images are loaded
            document.querySelector('.sect_main').style.opacity = '1';
        }
    }

    // Attach load event to each image
    for (let i = 0; i < images.length; i++) {
        if (images[i].complete) {
            checkAllImagesLoaded();
        } else {
            images[i].addEventListener('load', checkAllImagesLoaded);
        }
    }
});