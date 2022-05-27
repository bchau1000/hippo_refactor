class History {
    constructor() {}
}

export class Session {
    constructor(
        public history: History = new History()
    ) {}
}