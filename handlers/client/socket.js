var source = new EventSource("http://localhost:1234/sse/signal");
source.onmessage = function (event) {
    console.log(event)
    var message = event.data
    console.log(message)
  /*  document.querySelector('#qr').innerHTML = "";
    console.log(event)
    var message = event.data
    
   // const data = JSON.parse(event.data).data;
   // alert(data)
    //console.log(message)
    //const obj = JSON.parse(message);
    //console.log(obj)
    msg = document.querySelector('#message')
    msg.innerHTML = message.replace(/\"/g, "")+'<br>'+msg.innerHTML;
    if (new String(message).valueOf() == "success" || new String(message).valueOf() == "timeout/Refreshing"
        || new String(message).valueOf() == "Already logged") {
    } else {
     /  var qrcode = new QRCode("qr", {
            text: message,
            width: 128,
            height: 128,
            colorDark : "#000000",
            colorLight : "#ffffff",
            correctLevel : QRCode.CorrectLevel.M
        }); *
    } */
    const obj = JSON.parse(message)
    const Details = document.createElement("details");
    const Summary = document.createElement("summary");

    incomingMsg = obj.MessageText.replace(/\"/g, "")
    // const result = 10 > 5 ? 'yes' : 'no';
    const from = obj.Sender === obj.Group ? obj.Sender : obj.Sender + " / " + obj.Group;
    const msg = obj.MessageCaption === "" ? 
    incomingMsg : 
    incomingMsg +'<br>'+ "Caption: " + obj.MessageCaption;

    const body = obj.Uri === "" ? msg : msg + "<a href='" + obj.Uri + "' target='_blank'>Open attachment</a>"

    Summary.innerHTML = "From: "+ '<code>' + obj.Name + '</code>' + " Number: " + from + " By: " + obj.Time
    Details.setAttribute("type", obj.MessageType)
    Details.setAttribute("url", obj.Uri)
    Details.setAttribute("id", obj.MessageID)
    const Content = document.createElement("p");
    Content.innerHTML = "Text: " + body
    Details.appendChild(Summary);
    Details.appendChild(Content);
    console.log(Details)
    document.body.insertBefore(Details, document.body.children[1]);

    let ds= [...document.querySelectorAll('details')];
    ds.forEach(d=>d.addEventListener('click',e=>e.shiftKey||ds.filter(i=>i!=d).forEach(i=>i.removeAttribute('open'))))

    // <a href='%v' target='_blank'>Open</a>"
  //  Details.insertBefore(newDiv, currentDiv);

   // document.body.insertBefore(newDiv, currentDiv);
}