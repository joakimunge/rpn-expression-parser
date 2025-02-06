type Arithmetic = (a: number, b: number) => number;

interface Operators {
  [ops: string]: Arithmetic;
}

const operators: Operators = {
  '+': (a: number, b: number) => a + b,
  '-': (a: number, b: number) => a - b,
  '*': (a: number, b: number) => a * b,
  '/': (a: number, b: number) => a / b,
  '%': (a: number, b: number) => a % b,
  '^': (a: number, b: number) => a ** b,
};

function RPN(expression: string): number {
  let stack: Array<number> = [];

  for (let i = 0; i < expression.length; i++) {
    const item = expression[i];

    if (item === ' ') {
      continue;
    }

    const itemAsNum = Number(item);

    if (!isNaN(Number(item))) {
      stack.push(itemAsNum);
      continue;
    }

    if (item in operators) {
      const arithmetic = operators[item];
      const val1 = stack.shift();
      const val2 = stack.shift();
      if (!val1 || !val2) {
        throw new Error('Invalid RPN expression');
      }
      const result = arithmetic(val1, val2);
      stack.push(result);
    } else {
      throw new Error('Invalid expression');
    }
  }
  return stack.shift() ?? 0;
}
