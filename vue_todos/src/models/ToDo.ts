class ToDo {
  id: number
  title: string
  description?: string
  completed: boolean


  constructor() {
    this.id = 0
    this.title = ''
    this.description = ''
    this.completed = false
  }
}

export default ToDo
