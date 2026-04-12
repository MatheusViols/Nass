function displayMessage(message, successful) {
	let saveMessageBox = document.getElementById("save-message-box");
	let divMessage = document.createElement('div');
	let pMessage = document.createElement('p');

	if (successful) {
		divMessage.setAttribute('id', 'succeed-message');
	} else {
		divMessage.setAttribute('id', 'error-message');
	}

	pMessage.textContent = message;
	divMessage.appendChild(pMessage);
	saveMessageBox.appendChild(divMessage);

}

function removeMessage() {
	let saveMessageBox = document.getElementById("save-message-box");
	saveMessageBox.removeChild(saveMessageBox.firstChild);
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
