export abstract class Model {
    private id: number;

    constructor(id: number) {
        this.id = id;
    }

    public getId(): number {
        return this.id;
    }

    public toString(): string {
        return `Model{id: ${this.id}}`;
    }
}