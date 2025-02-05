from leetscrape import GetQuestion
from bs4 import BeautifulSoup
from concurrent import futures
import re
import bind_pb2
import bind_pb2_grpc
import grpc
def clean_text(html):
    return BeautifulSoup(html, "html.parser").get_text()
class StatementServicer(bind_pb2_grpc.StatementServiceServicer):
    def GetStatement(self,request,context):
        question = GetQuestion(titleSlug=request.title_slug).scrape()
        text = clean_text(question.Body)
        match = re.search(r"(.*?)\nExample 1:", text, re.DOTALL)
        Statement = match.group(1) if match else "Statement not found."
        Statement=re.sub(r'\s+', ' ', Statement).strip()
        hints=""
        for hint in question.Hints:
            temp=clean_text(hint)
            hints+=temp+"\n"
        result = f"""
Question : {Statement}
Hints : {hints}
        """
        return bind_pb2.ProblemStatement(statement=result)

def serve():
    server=grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    bind_pb2_grpc.add_StatementServiceServicer_to_server(StatementServicer(),server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
    
if __name__=='__main__':
    serve()