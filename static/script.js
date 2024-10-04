function onInputChange() {
    const textInput = document.getElementById("textInput");
    const text = textInput.value;
    const style = document.getElementById("style").value;
    const data = {text, style};
    const json = JSON.stringify(data);
    // Make AJAX request to Go server
    fetch(`/ascii-art`,{
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: json,
  
    })
      .then(response => response.text())
      .then(data => {
        const asciiArt = data;
        document.getElementById("result").textContent = asciiArt;
      })
      .catch(error => console.error(error));
  }