function setModalTitle(title) {
    let modalTitle = document.querySelector("#todo_create_update_modal .modal-title")
    modalTitle.textContent = title
}

function resetInput() {
    let titleInput = document.querySelector("#title_input")
    titleInput.value = ""
}

function createUpdateTask(itemID, taskTitle, method) {
    setModalTitle(taskTitle)
    let request_to_server_btn_create_update = document.querySelector('#request_to_server_btn_create_update')
    request_to_server_btn_create_update.setAttribute("onclick", `requestToServerCreateUpdate('${itemID}', '${method}')`)
}

function deleteTask(itemID) {
    let request_to_server_btn_delete = document.querySelector('#request_to_server_btn_delete')
    request_to_server_btn_delete.setAttribute("onclick", `requestToServerDelete('${itemID}')`)
}

async function requestToServerCreateUpdate(itemID, method) {
    let debugUrl = 'http://localhost:8000/todos/' + itemID
    let releaseUrl = 'http://localhost/todos/' + itemID
    let title = document.querySelector("#title_input").value
    if (title) {
        let url = releaseUrl
        let data = {
            "title": title,
        }
        let header = {
            "Accept": "application/json",
            "Content-Type": "application/json",
        }
        let init = {
            method: method,
            headers: header,
            body: JSON.stringify(data),
            mode: "cors",
        }
        let response = await fetch(url, init)
        await response.json().then(data => {
            if (response.status === 200 || response.status === 201) {
                window.location.reload()
            }
        })
    }
}

async function requestToServerDelete(itemID) {
    let debugUrl = 'http://localhost:8000/todos/' + itemID
    let releaseUrl = 'http://localhost/todos/' + itemID
    let url = releaseUrl
    let header = {
        "Accept": "application/json",
        "Content-Type": "application/json",
    }
    let init = {
        method: "DELETE",
        headers: header,
        mode: "cors",
    }
    let response = await fetch(url, init)
    await response.json().then(data => {
        if (response.status === 200) {
            window.location.href = "/todos"
        }
    })
}