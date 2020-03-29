const btnCopy = document.querySelector(".copy")


btnCopy.onclick = () => {
    const el = document.createElement('textarea');
    el.value = document.querySelector('#code').value;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
}
