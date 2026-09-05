export const snakeToTitleCase = (value: string): string =>
  value
    ?.replace(/_/g, " ")
    .replace(/\b\w/g, (char) => char.toUpperCase()) ?? "";

export const snakeToLowerCaseSentence = (value: string): string =>
  value?.replace(/_/g, " ").toLowerCase() ?? "";