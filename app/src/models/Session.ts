class History {
    constructor({
        
    }) {}
}

export class Session {
    private history: History;

    constructor({ history }:{ history: History }) {
        this.history = history;
    }

    public getHistory(): History {
        return this.history;
    }
}