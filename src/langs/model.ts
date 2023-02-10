export interface dattType {
  main: {
    title: string;
    description: string;
    home: {
      get_result: string;
      made_by: string;
    };
    result: {
      result_tab: {
        title: string;
        summary: string;
        unattempted_ques: string;
        attempted_ques: string;
        correct_ques: string;
        incorrect_ques: string;
        total_score: string;
        detailed_report: string;
        table: {
          given_answer: string;
          correct_answer: string;
          question_id: string;
          correct_answer_id: string;
          given_answer_id: string;
        };
      };
      questions_tab: {
        title: string;
        correct_answer: string;
        marked_answer: string;
      };
    };
  };
}