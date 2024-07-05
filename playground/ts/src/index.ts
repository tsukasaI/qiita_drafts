const sampleClassDecorator = (constructor: Function) => {
    Object.seal(constructor);
    Object.seal(constructor.prototype);
    console.log('sampleClassDecorator called.');
}

const sampleMethodDecorator = (value: boolean) => {
    return (target: any, propertyKey: string, descriptor: PropertyDescriptor) => {
        descriptor.enumerable = value;
       console.log('sampleMethodDecorator called.');
   }
}


@sampleClassDecorator
class SampleDecoratedClass {
    constructor() {
        console.log('SampleDecoratedClass constructor called.');
    }
}

class SampleMethodDecoratedClass {
    @sampleMethodDecorator(false)
    sampleMethod() {
        console.log('SampleMethodDecoratedClass sampleMethod called.');
    }
}


const main = () => {
    const sampleClassInstance = new SampleDecoratedClass();

    const sampleMethodDecoratedClassInstance = new SampleMethodDecoratedClass();
    sampleMethodDecoratedClassInstance.sampleMethod();
}
