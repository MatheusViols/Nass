function displayMessage(message, successful) {
	let saveMessageBox = document.getElementById("save-message-box");
	let pMessage = document.createElement('p');
	pMessage.id = "message";

	if (successful) {
		pMessage.style.color = "green";
	} else {
		pMessage.style.color = "red";
	}

	pMessage.textContent = message;
	saveMessageBox.appendChild(pMessage);
}

function removeMessage() {
	let saveMessageBox = document.getElementById("save-message-box");
	let pMessage = document.getElementById("message");

	saveMessageBox.removeChild(pMessage);
}


function getSaveRequest() {
	let editorForm = document.getElementById("editor");

	let data = new FormData(editorForm);
	data = Object.fromEntries(data.entries());

	return {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(data),
	};        
}

async function saveNote(event) {
	event.preventDefault();

	let req = getSaveRequest();

	let res = await fetch("/save/", req).catch(error => console.log(error.message));

	let resData = await res.json();

	displayMessage(resData.Message, res.ok);

	setTimeout(() => {
		removeMessage()
	}, 3000)
}

let saveButton = document.getElementById("save-button");
saveButton.addEventListener("click", saveNote);
