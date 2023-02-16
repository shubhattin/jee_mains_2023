export interface dattType {
  main: {
    title: string;
    description: string;
    home: {
      get_result: string;
      get_data_result: string;
      DOB: string;
      appl_numb: string;
      get_sample_result: string;
      verify: {
        appl_numb_verify: string;
        dob_verify: string;
      };
      made_by: string;
      page_views: string;
      result_views: string;
      data_page: {
        first_time_info: string;
        response_page_data: string;
        response_page_data_msg: string;
        answer_key_data: string;
        answer_key_data_msg: string;
        submit_btn: string;
        error_msg: string;
        notes: {
          get_html_page: string;
          on_mobile: string;
          on_mobile_msg: string;
          on_computer: string;
          on_computer_msg: string;
          other_notes: [string, string];
        };
      };
    };
    result: {
      sample_result_msg: string;
      result_tab: {
        title: string;
        summary: string;
        unattempted_ques: string;
        attempted_ques: string;
        correct_ques: string;
        incorrect_ques: string;
        total_score: string;
        detailed_report: string;
        subject_wise: {
          title: string;
          subjects: [string, string, string];
          table: {
            subject: string;
            correct: string;
            incorrect: string;
            unattempted: string;
            score: string;
          };
        };
        table: {
          given_answer: string;
          correct_answer: string;
          question_id: string;
          correct_answer_id: string;
          given_answer_id: string;
        };
      };
      questions_tab: {
        desktop_mode_better: string;
        title: string;
        correct_answer: string;
        marked_answer: string;
      };
    };
  };
}
