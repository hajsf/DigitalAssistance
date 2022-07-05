var source = new EventSource("http://localhost:1234/sse/signal");
source.onmessage = function (event) {
    document.querySelector('#qr').innerHTML = "";
    var message = event.data
    document.querySelector('#message').innerHTML = message;
    if (new String(message).valueOf() == "success" || new String(message).valueOf() == "timeout/Refreshing"
        || new String(message).valueOf() == "Already logged") {
    } else {
        var qrcode = new QRCode("qr", {
            text: message,
            width: 128,
            height: 128,
            colorDark : "#000000",
            colorLight : "#ffffff",
            correctLevel : QRCode.CorrectLevel.M
        });
    }
}