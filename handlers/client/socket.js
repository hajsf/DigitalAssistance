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
    const Content = document.createElement("p");
    const Attachment = document.createElement(null);

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
    
    Content.innerHTML = "Text: " + body

    if(obj.MessageType === "image") {
      Attachment.innerHTML = '<br>' + "<img src='"+obj.Uri+"' alt='Message attachment' width: 100%;' height='600'>"
    } else if(obj.MessageType === "audio"){
    //  console.log(`Audio source: ${obj.Uri}`)
      Attachment.innerHTML = '<br>' + "<audio controls> "+
        "<source src='"+obj.Uri+"' type='audio/ogg'>" +
        "<source src='"+obj.Uri+"' type='audio/mpeg'>" +
        "<source src='"+obj.Uri+"' type='audio/oga'>" +
     //   "<source src='test.oga' type='audio/ogg; codecs=`vorbis`'></source>" +
        "Your browser does not support the audio element." +
      "</audio>"
    } else if(obj.MessageType === "video"){
      Attachment.innerHTML = '<br>' + "<video width: 100%;' height='600' controls> " +
        "<source src='"+obj.Uri+"' type='video/ogg'>" +
        "<source src='"+obj.Uri+"' type='video/mp4'>" +
        "<source src='"+obj.Uri+"' type='video/m4v'>" +
        "Your browser does not support the video tag." +
      "</video>"
    } else if(obj.MessageType === "document"){
      Attachment.innerHTML = '<br>' +  "<embed id='pdf' type='application/pdf'" +
      "src='"+obj.Uri+"' style='width: 100%;' height='600'>"
    }

    Content.appendChild(Attachment);

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