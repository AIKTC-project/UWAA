export const showPassword = () => {
    document.getElementsByClassName('eye')[0].classList.toggle('fill')
    const t = document.getElementsByClassName('eye')[0].parentNode.children[0].getAttribute('type')
    if (t === 'password') document.getElementsByClassName('eye')[0].parentNode.children[0].setAttribute('type', 'text')
    else if (t === 'text') document.getElementsByClassName('eye')[0].parentNode.children[0].setAttribute('type', 'password')
  }