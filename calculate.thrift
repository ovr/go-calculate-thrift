namespace php tutorial
namespace cpp tutorial
namespace go tutorial

enum Operation {
  ADD = 1,
  SUBTRACT = 2,
  MULTIPLY = 3,
  DIVIDE = 4
}

struct Work {
  1: i32 num1 = 0,
  2: i32 num2,
  3: Operation op,
  4: optional string comment,
}

exception InvalidOperation {
  1: i32 whatOp,
  2: string why
}

service Calculator {
   void ping(),

   i64 plus(1:i32 num1, 2:i32 num2),

   i32 minus(1:i32 num1, 2:i32 num2),

   i64 mul(1:i32 num1, 2:i32 num2),

   i32 div(1:i32 num1, 2:i32 num2),

   double mod(1:double num1, 2:double num2),

   i32 pow(1:i32 num1, 2:i32 num2),

   i32 calculate(1:i32 logid, 2:Work w) throws (1:InvalidOperation ouch),
}

