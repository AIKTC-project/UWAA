export const makeForm = (key: Array<string>, value: Array<string>): FormData => {
  const b = new FormData()
  for (let i = 0; i < key.length; i++) {
    b.set(key[i], value[i])
  }
  return b
}