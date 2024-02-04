// Reset Form & Response Input Textarea
document.getElementById('aesEncryptForm').reset();
document.getElementById('aesDecryptForm').reset();
document.getElementById('jsonEscapedForm').reset();
document.getElementById('jsonBase64ToImageForm').reset();
document.getElementById('jsonCamelToSnakeForm').reset();
document.getElementById('jsonImageToBase64').reset();
document.getElementById('jsonSnakeToCamelForm').reset();
document.getElementById('specialistDeleteForm').reset();
document.getElementById('specialistMakeForm').reset();

document.getElementById('responseEncryptText').value = '';
document.getElementById('responseDecryptText').value = '';
document.getElementById('responseJsonEscapedText').value = '';
document.getElementById('responseJsonCamelToSnakeText').value = '';
document.getElementById('responseJsonImageToBase64Text').value = '';
document.getElementById('responseJsonSnakeToCamelText').value = '';

// document.getElementById('responseSpecialistMake').value = '';

function generateGUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        let r = Math.random() * 16 | 0, v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

// Action Submit Aes Encrypt & Decrypt
function submitAesEncryptAndDecryptForm(formId, responseText) {
    this.event.preventDefault();
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    const data = {};

    formData.forEach((value, key) => {
        data[key] = value;
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method, headers: {
            'Content-Type': 'application/json',
            'x-correlation-id': guid
        }, body: JSON.stringify(data),
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById(responseText).value = result;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// Action Submit Specialist Make
function submitSpecialistMakeForm(formId, responseText) {
    this.event.preventDefault();
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    const data = new FormData();

    formData.forEach((value, key) => {
        data.append(key, value)
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method,
        headers: {
            'x-correlation-id': guid
        }, body: data,
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById(responseText).value = result;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// Action Submit Specialist Delete
function submitSpecialistDeleteForm(formId) {
    this.event.preventDefault();
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    const data = new FormData();

    formData.forEach((value, key) => {
        data.append(key, value)
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method,
        headers: {
            'x-correlation-id': guid
        }, body: data,
    })
        .then(response => response.blob())
        .then(blob => {
            const url = window.URL.createObjectURL(blob);

            // Create a temporary anchor element
            const a = document.createElement('a');
            a.href = url;
            a.download = 'EIM_EDGE_BLACKLIST_' + getCurrentYearMonthDay() + '.txt';
            document.body.appendChild(a);

            // Trigger a click on the anchor to start the download
            a.click();

            // Remove the temporary anchor element
            document.body.removeChild(a);

            createCtrlFileSpecialist()
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


function getCurrentYearMonthDay() {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = String(currentDate.getMonth() + 1).padStart(2, '0'); // Months are zero-based
    const day = String(currentDate.getDate()).padStart(2, '0');
    return `${year}${month}${day}`;
}

function createCtrlFileSpecialist() {
    // Create an empty Blob
    const blob = new Blob([''], {type: 'text/plain'});

    // Create a download link
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = 'CTRL_EIM_EDGE_BLACKLIST_' + getCurrentYearMonthDay() + '.txt';

    // Trigger the download
    link.click();

    // Cleanup
    URL.revokeObjectURL(link.href);
}

// Action Submit Json Transform "camel-to-snake","snake-to-camel","json-escaped"
function submitJsonTransformForm(formId, responseText) {
    this.event.preventDefault();
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    let dataSend = '';

    formData.forEach((value) => {
        dataSend = value;
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method, headers: {
            'Content-Type': 'application/json',
            'x-correlation-id': guid
        }, body: dataSend,
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById(responseText).value = result;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// Action Submit Base64 To Image
function submitBase64ToImageForm(formId) {
    this.event.preventDefault(); // Prevent the default form submission
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    let dataSend = '';

    formData.forEach((value) => {
        dataSend = value;
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method, headers: {
            'Content-Type': 'application/json',
            'x-correlation-id': guid
        }, body: dataSend,
    }).then(response => response.blob())
        .then(blob => {
            const url = window.URL.createObjectURL(blob);

            // Create a temporary anchor element
            const a = document.createElement('a');
            a.href = url;
            a.download = 'image_' + new Date().getTime(); // You can customize the filename here
            document.body.appendChild(a);

            // Trigger a click on the anchor to start the download
            a.click();

            // Remove the temporary anchor element
            document.body.removeChild(a);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


// Action Submit Image To Base64
function submitImageToBase64Form(formId, responseText) {
    this.event.preventDefault();
    const form = document.getElementById(formId);
    const formData = new FormData(form);
    let dataSend = '';

    formData.forEach((value) => {
        dataSend = value;
    });

    let guid = generateGUID();
    fetch(form.action, {
        method: form.method, headers: {
            'Content-Type': 'application/json',
            'x-correlation-id': guid
        }, body: dataSend,
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById(responseText).value = result;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


function formatAsJson(responseTextId) {
    try {
        const responseText = document.getElementById(responseTextId).value;
        const json = JSON.parse(responseText);
        document.getElementById(responseTextId).value = JSON.stringify(json, null, 2);
    } catch (error) {
        console.error('Error formatting as JSON:', error);
    }
}

function resetFormAndResponseTextarea(formId, textResponseId) {
    document.getElementById(formId).reset();
    if (textResponseId !== null) {
        document.getElementById(textResponseId).value = '';
    }

}

function copyToClipboard(copyTextareaId) {
    const copyTextarea = document.getElementById(copyTextareaId);
    copyTextarea.select();
    navigator.clipboard.writeText(copyTextarea.value)
        .then(() => {
            console.log('Text successfully copied to clipboard');
        })
        .catch(err => {
            console.error('Unable to copy text to clipboard', err);
        });
}

function changeTab(tabId, divFormId) {

    const ulElement = document.getElementById('tabList');
    const liElements = ulElement.getElementsByTagName('li');

    for (let i = 0; i < liElements.length; i++) {
        const aElement = liElements[i].getElementsByTagName('a')[0].id;

        if (aElement !== tabId) {
            const elementInactive = document.getElementById(aElement);
            elementInactive.classList.remove(...elementInactive.classList);
            elementInactive.classList.add('inline-flex', 'items-center', 'px-4', 'py-3', 'rounded-lg', 'hover:text-gray-900', 'bg-gray-50', 'hover:bg-gray-100', 'w-full', 'dark:bg-gray-800', 'dark:hover:bg-gray-700', 'dark:hover:text-white');
        } else {
            const elementActive = document.getElementById(tabId);
            elementActive.classList.remove(...elementActive.classList);
            elementActive.classList.add('inline-flex', 'items-center', 'px-4', 'py-3', 'text-white', 'bg-blue-700', 'rounded-lg', 'active', 'w-full', 'dark:bg-blue-600');
        }
    }

    const divElements = document.querySelectorAll('div[id^="tab"]');
    divElements.forEach((div) => {
        let elementActive = document.getElementById(div.id);
        if (div.id === divFormId) {
            elementActive.classList.remove('hidden');
        } else {
            elementActive.classList.add('hidden');
        }
    });
}