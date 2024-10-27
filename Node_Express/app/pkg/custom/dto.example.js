class LastUpdateDTO {
  constructor(element) {
    this.things = {};
    Object.assign(this.things, element);
    this.things.updated = new Date();
  }
}

export default LastUpdateDTO;