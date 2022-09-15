export const toggleRed = async (elements: Array<HTMLInputElement> | Array<EventTarget>, after4s?: () => void, forceToggle ?:boolean) => {
  for (let i = 0; i < elements.length; i++) {
    if (elements[i].value.length === 0 || forceToggle) {
      elements[i].setAttribute('style', 'outline: 1px solid #d00; box-shadow: 0 0 0.5rem #d00;')
      setTimeout(() => {
        elements[i].removeAttribute('style')
      }, 4000)
    }
    setTimeout(after4s, 4000)
  }
}