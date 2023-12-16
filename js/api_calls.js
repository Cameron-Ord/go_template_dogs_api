
const socket = new WebSocket("ws://localhost:8080/ws");
socket.addEventListener("open", (event) => {
    console.log("WebSocket connection opened:", event);
});
socket.addEventListener("message", (event) => {
    
    const message = JSON.parse(event.data);
    if (message.length > 0){
        console.log("greater")
        const dog_div = document.createElement("div");
        for(let i = 0; i < message.length; i++){
            const img_element = document.createElement("img");
            img_element.src = message[i]["url"]
            img_element.alt = "Doggy";
            img_element.style.width = "250px"
            img_element.style.height = "250px"
            img_element.style.objectFit = "cover"
            dog_div.appendChild(img_element);
        }

        const body = document.querySelector('body');
        body.appendChild(dog_div);
    }


});

socket.addEventListener("close", (event) => {
    console.log("WebSocket connection closed:", event);
});

let button = document.getElementById('click_me');
button.addEventListener("click", () => {
    console.log('clicked')
    socket.send("Give me DOGS");
});
